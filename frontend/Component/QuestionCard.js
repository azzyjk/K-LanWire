import React from "react";
import { View, Text, StyleSheet, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";

export default function QuestionCard({
  navigation,
  num,
  question,
  solved,
  answers,
  questionYear,
  questionDepartment,
}) {
  return (
    <View style={styles.inner}>
      <View style={styles.question}>
        <Text
          style={{ color: "gray" }}
        >{` < ${questionYear} ${questionDepartment} >`}</Text>
        <Text />
        <Text
          style={styles.questionFont}
          numberOfLines={1}
          ellipsizeMode="tail"
        >
          Q:{" "}
          {question.length < 30
            ? `${question}`
            : `${question.substring(0, 20)}...`}
          {solved == true ? (
            <Ionicons name="md-checkmark" size={20} color="green" />
          ) : (
            <Text />
          )}
        </Text>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    borderTopLeftRadius: 20,
    borderTopRightRadius: 20,
  },
  inner: {
    flex: 1,
    paddingLeft: 10,
    padding: 20,
    borderBottomWidth: 1,
    borderBottomColor: "gray",
  },
  question: {
    marginLeft: 10,
    marginRight: 20,
  },
  questionFont: {
    fontSize: 20,
  },
});
