import React from "react";
import {
  View,
  Text,
  ScrollView,
  SafeAreaView,
  StyleSheet,
  KeyboardAvoidingView,
  Platform,
} from "react-native";

import Answer from "../Component/Answer";

export default function QuestionScreen({ route, navigation }) {
  const {
    num,
    question,
    solved,
    answers,
    questionDepartment,
    questionYear,
  } = route.params;
  console.log(
    `${num}\n ${question}\n ${solved}\n ${answers}\n ${questionDepartment}\n ${questionYear}`
  );
  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      style={styles.container}
    >
      <View style={{ flex: 5, justifyContent: "space-around" }}>
        <SafeAreaView style={styles.question}>
          <ScrollView>
            <View style={styles.card}>
              <Text style={{ fontSize: 40, fontWeight: "900" }}>Q</Text>
              <Text></Text>
              <Text style={{ color: "gray" }}>
                {`  < ${questionYear} ${questionDepartment} >`}
              </Text>
              <View style={{ margin: 20 }}>
                <Text style={styles.questionFont}> {question}</Text>
                <Text></Text>
              </View>
            </View>
            {solved == true ? (
              Object.keys(answers).map((idx) => {
                return (
                  <View style={styles.card} key={idx}>
                    <View style={{ margin: 15 }}>
                      <Text style={{ color: "gray" }}>
                        {`< ${answers[idx]["AnswerYear"]} ${answers[idx]["AnswerDepartment"]} >`}
                      </Text>
                      <Text style={{ color: "gray" }}>{""}</Text>
                      <Text style={{ fontSize: 15 }}>
                        {answers[idx]["Answer"]}
                      </Text>
                    </View>
                  </View>
                );
              })
            ) : (
              <View></View>
            )}
          </ScrollView>
        </SafeAreaView>
      </View>
      <View style={styles.answer}>
        <Answer num={num} navigation={navigation} />
      </View>
    </KeyboardAvoidingView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "white",
  },
  inner: {
    flex: 1,
    justifyContent: "space-around",
  },
  question: {
    flex: 8,
  },
  questionFont: {
    fontSize: 20,
  },
  answer: {
    flex: 2,
    paddingHorizontal: 24,
  },
  card: {
    justifyContent: "center",
    flex: 1,
    borderColor: "gray",
    borderBottomWidth: 1,
  },
});
