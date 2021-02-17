import React, { useCallback, useState } from "react";
import { SafeAreaView, ScrollView, View } from "react-native";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { useFocusEffect } from "@react-navigation/native";
import axios from "axios";

export default function QnAScreen({ navigation }) {
  const [refreshing, setRefreshing] = useState(false);
  const [quetions, setQuestions] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const _getQuestions = async () => {
    const department = await AsyncStorage.getItem("department");
    const sendData = {
      headers: {
        Accept: "utf-8",
        "content-type": "application/x-www-form-urlencoded; charset=utf-8",
      },
      data: JSON.stringify({
        학과: department,
      }),
      method: "POST",
      url: "http://119.192.41.219:9184/qna_enter",
    };
    axios(sendData).then((res) => {
      setQuestions(res.data.QuestionList);
    });
  };

  useFocusEffect(
    useCallback(() => {
      (async () => {
        await _getQuestions();
        setIsLoading(true);
      })();
    })
  );

  if (isLoading == false) return <View />;
  else {
    return (
      <SafeAreaView>
        <ScrollView></ScrollView>
      </SafeAreaView>
    );
  }
}
