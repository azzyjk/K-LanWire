import React, { useState } from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import AsyncStorage from "@react-native-async-storage/async-storage";

import Logo from "../Component/Logo";
import DepartmentPicker from "../Component/DepartmentPicker";
import EntranceYearPicker from "../Component/EntranceYearPicker";
import JobPicker from "../Component/JobPicker";

export default function InitScreen({ haveInfo, setHaveInfo }) {
  const [department, setDepartment] = useState("컴퓨터공학부");
  const [entranceYear, setEntranceYear] = useState("2021");
  const [job, setJob] = useState("Professor");

  const _onPressButton = async () => {
    await AsyncStorage.setItem("department", department);
    await AsyncStorage.setItem("job", job);
    if (job == "Professor") await AsyncStorage.setItem("entranceYear", "0");
    else await AsyncStorage.setItem("entranceYear", entranceYear);
    setHaveInfo(true);
  };

  return (
    <View style={styles.container}>
      <View style={styles.logoComponent}>
        <Logo />
      </View>
      <View style={styles.pickerComponent}>
        <DepartmentPicker
          department={department}
          setDepartment={setDepartment}
        />
      </View>
      <JobPicker job={job} setJob={setJob} />
      <View style={styles.pickerComponent}>
        <EntranceYearPicker
          entranceYear={entranceYear}
          setEntranceYear={setEntranceYear}
          job={job}
        />
      </View>
      <View style={styles.buttonComponent}>
        <TouchableOpacity onPress={_onPressButton}>
          <Text style={styles.buttonFont}> 시작하기 </Text>
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
    paddingTop: 30,
  },
  pickerComponent: {
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
