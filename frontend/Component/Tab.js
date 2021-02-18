import React from "react";
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  Dimensions,
  Platform,
  StatusBar,
} from "react-native";
import { Picker } from "@react-native-picker/picker";
import { Ionicons } from "@expo/vector-icons";

const width = Dimensions.get("window").width;
const height = Dimensions.get("window").height;

export default function Tab({ navigation, name }) {
  return (
    <View style={styles.container}>
      <TouchableOpacity style={{ flex: 1 }} onPress={() => navigation.goBack()}>
        <Ionicons name="ios-arrow-back" size={25} color="#3399ff" />
      </TouchableOpacity>
      <View style={{ flex: 1, alignItems: "center" }}>
        <Text style={styles.titleFont}> {name}</Text>
      </View>
      <View style={{ flex: 1 }}></View>
    </View>
  );
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: "row",
    margin: 10,
    paddingTop: 40,
  },
  titleFont: {
    fontSize: 20,
    fontWeight: "600",
  },
});
