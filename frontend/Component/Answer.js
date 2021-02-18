import React, { useEffect, useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  TextInput,
  Alert,
} from "react-native";
import axios from "axios";
import AsyncStorage from "@react-native-async-storage/async-storage";

export default function Answer({ num, navigation }) {
  const [value, onChnageText] = useState("");
  const [department, setDepartment] = useState();
  const [entranceYear, setEntranceYear] = useState();
  const [job, setJob] = useState();

  useEffect(() => {
    (async () => {
      setDepartment(await AsyncStorage.getItem("department"));
      setEntranceYear(await AsyncStorage.getItem("entranceYear"));
      if ((await AsyncStorage.getItem("job")) == "Professor") {
        setJob("1");
      } else {
        setJob("0");
      }
    })();
  }, []);

  var sendData = {
    headers: {
      Accept: "utf-8",
      "content-type": "application/x-www-form-urlencoded; charset=utf-8",
    },
    data: JSON.stringify({
      questionNum: String(num),
      answerEntranceYear: entranceYear,
      answerDepartment: department,
      answer: value,
      rank: job,
    }),
    method: "POST",
    url: "http://119.192.41.219:9184/qna_answer",
  };

  const _onPressButton = () => {
    axios(sendData)
      .then((res) => {
        Alert.alert("답변이 등록되었습니다.");
        navigation.goBack();
      })
      .catch((err) => {
        Alert.alert(err.response.data["message"]);
      });
  };

  return (
    <View style={{ flex: 1 }}>
      <TextInput
        style={styles.textInput}
        onChangeText={(text) => onChnageText(text)}
        value={value}
        placeholder="답변을 입력해주세요."
      />

      <TouchableOpacity style={styles.buttonComponent} onPress={_onPressButton}>
        <Text style={styles.buttonFont}>답변하기</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  textInput: {
    height: 40,
    borderColor: "#000000",
    borderBottomWidth: 1,
    marginBottom: 36,
  },
  buttonComponent: {
    flex: 1,
    alignItems: "center",
    marginTop: 40,
  },
  buttonFont: {
    fontSize: 30,
    fontWeight: "900",
    color: "gray",
  },
});
