import React from "react";
import { View, StyleSheet } from "react-native";

export default function MainScreen({ navigation }) {
  return <View style={styles.container}></View>;
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "white",
  },
});
