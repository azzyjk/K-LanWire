import React, { useState } from "react";
import { View, StyleSheet, Text, Alert } from "react-native";

import Logo from "../Component/Logo";
import DepartmentPicker from "../Component/DepartmentPicker";
import EntranceYearPicker from "../Component/EntranceYearPicker";
import JobPicker from "..//Component/JobPicker";
import { TouchableOpacity } from "react-native-gesture-handler";
import AsyncStorage from "@react-native-async-storage/async-storage";

export default function SettingScreen({ navigation }) {
  const [department, setDepartment] = useState("컴퓨터공학부");
  const [entranceYear, setEntranceYear] = useState("2021");
  const [job, setJob] = useState("Professor");

  const _onPressButton = async () => {
    await AsyncStorage.setItem("department", department);
    await AsyncStorage.setItem("entranceYear", entranceYear);
    await AsyncStorage.setItem("job", job);
    Alert.alert("수정되었습니다.");
    navigation.goBack();
  };

  return (
    <View style={styles.container}>
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