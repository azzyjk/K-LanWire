import React, { useState } from "react";
import { View, StyleSheet, Text, Alert } from "react-native";

import { TouchableOpacity } from "react-native-gesture-handler";
import AsyncStorage from "@react-native-async-storage/async-storage";

import Logo from "../Component/Logo";
import DepartmentPicker from "../Component/DepartmentPicker";
import EntranceYearPicker from "../Component/EntranceYearPicker";
import JobPicker from "..//Component/JobPicker";
import Tab from "../Component/Tab";

export default function SettingScreen({ navigation }) {
  const [department, setDepartment] = useState("컴퓨터공학부");
  const [entranceYear, setEntranceYear] = useState("2019");
  const [job, setJob] = useState("Professor");

  const _onPressButton = async () => {
    await AsyncStorage.setItem("department", department);
    await AsyncStorage.setItem("job", job);
    if (job == "Professor") await AsyncStorage.setItem("entranceYear", "2000");
    else await AsyncStorage.setItem("entranceYear", entranceYear);
    Alert.alert("수정되었습니다.");
    navigation.navigate("Chat",{data : "data"});
  };

  return (
    <View style={styles.container}>
      <Tab navigation={navigation} name={"Setting"} />
      <View style={styles.logoComponent}>
        <Logo />
      </View>
      <View style={styles.picker}>
        <DepartmentPicker
          department={department}
          setDepartment={setDepartment}
        />
      </View>
      <JobPicker job={job} setJob={setJob} />
      <View style={styles.picker}>
        <EntranceYearPicker
          entranceYear={entranceYear}
          setEntranceYear={setEntranceYear}
          job={job}
        />
      </View>
      <View style={styles.buttonComponent}>
        <TouchableOpacity onPress={_onPressButton}>
          <Text style={styles.buttonFont}> 수정하기 </Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
  logoComponent: {
    flex: 2,
    justifyContent: "center",
  },
  picker: {
    flex: 2,
    justifyContent: "center",
  },
  buttonComponent: {
    flex: 2,
    justifyContent: "center",
  },
  buttonFont: {
    fontSize: 40,
  },
});
