import React from "react";
import { TouchableOpacity, View, StyleSheet, Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";

import Logo from "../Component/Logo";

export default function MainScreen({ navigation }) {
  return (
    <View style={styles.container}>
      <TouchableOpacity
        style={styles.settingButton}
        onPress={() => navigation.navigate("Setting")}
      >
        <Ionicons name="md-settings" size={32} color="black" />
      </TouchableOpacity>
      <View style={styles.logoComponent}>
        <Logo />
      </View>
      <TouchableOpacity style={styles.buttonComponent}>
        <Text style={styles.buttonFont}> 챗봇</Text>
      </TouchableOpacity>
      <TouchableOpacity style={styles.buttonComponent}>
        <Text style={styles.buttonFont}> Q & A 게시판</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "white",
  },
  settingButton: {
    alignItems: "flex-end",
    paddingRight: 20,
    paddingTop: 20,
  },
  logoComponent: {
    flex: 1,
    alignItems: "center",
  },
  buttonComponent: {
    flex: 1,
    alignItems: "center",
  },
  buttonFont: {
    fontSize: 40,
  },
});
