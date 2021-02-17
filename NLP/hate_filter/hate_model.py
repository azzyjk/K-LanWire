
# https://huggingface.co/transformers/installation.html
# pip install transformers

import tensorflow as tf
from transformers import ElectraTokenizer, TFElectraModel, logging
import numpy as np
import re
from hanspell import spell_checker

EMB_LEN = 76
DEFAULT_THRESH = 0.7

def load_model():
    electra_model = TFElectraModel.from_pretrained('monologg/koelectra-base-v3-discriminator', from_pt=True)

    input_toks = tf.keras.layers.Input(shape=(EMB_LEN,), name='toks', dtype='int32')
    input_masks = tf.keras.layers.Input(shape=(EMB_LEN,), name='masks', dtype='int32')

    electra_output = electra_model(input_toks, attention_mask=input_masks).last_hidden_state

    x = tf.keras.layers.GlobalAveragePooling1D()(electra_output)
    # x = tf.keras.layers.Dense(128, activation='relu')(x)
    # x = tf.keras.layers.Dropout(0.2)(x)
    # x = tf.keras.layers.Dense(32, activation='relu')(x)
    # x = tf.keras.layers.Dropout(0.2)(x)
    y = tf.keras.layers.Dense(2, activation='softmax')(x)

    model = tf.keras.models.Model(inputs=[input_toks, input_masks], outputs=y)
    model.load_weights('hate_weights/')

    return model

logging.set_verbosity_error()
tokenizer = ElectraTokenizer.from_pretrained('monologg/koelectra-base-v3-discriminator')
model = load_model()


def preprocess(texts):

    clean_texts = []

    for text in texts:
        cleaned = re.sub('[^가-힣\s]', '', text)
        cleaned = spell_checker.check(cleaned).checked
        clean_texts.append(cleaned)
    
    return clean_texts


# texts: string들이 담겨 있는 list
# thresh: softmax 분포의 악플 분류 기준치
#   ex) thresh가 0.7이면 softmax 분포가 [0.6, 0.4] 여도 악플로 분류, [0.8, 0.2] 는 정상으로 분류
#   (보니까 일반적인 답변은 보통 [0.95, 0.05] 이런식이기 때문에 0.7 정도로 잡아도 괜찮을거 같음)
def predict(texts, thresh=DEFAULT_THRESH):

    texts = preprocess(texts)

    test_X = tokenizer(texts, truncation=True, padding='max_length', max_length=EMB_LEN)
    test_toks = np.asarray(test_X['input_ids'])
    test_masks = np.asarray(test_X['attention_mask'])

    preds = model.predict(x={'toks':test_toks, 'masks':test_masks})
    labels = []

    for entry in preds:
        if entry[0] > thresh:
            labels.append(0)
        else:
            labels.append(1)
    
    return labels  # 0과 1이 담긴 list, 0: 정상 1: 악플


if __name__ == '__main__':
    examples = ['아니 이런걸 왜 모르지 이런건 고등학생도 알겠다 대학생 맞냐', 
                '신공학관 1층에도 있고 학생회관 2층에도 있음',
                '니가 먼저 직접 알아보고 여기다가 질문하라고']

    result = predict(examples)
    print(result) # [1, 0, 1]