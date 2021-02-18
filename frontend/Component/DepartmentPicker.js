import React from "react";
import { View, StyleSheet } from "react-native";
import { Picker } from "@react-native-picker/picker";

export default function DepartmentPicker({ department, setDepartment }) {
  const departments = ["컴퓨터공학부", "전기전자공학부", "기계항공공학부"];
  return (
    <View style={styles.container}>
      <Picker
        selectedValue={department}
        style={styles.picker}
        onValueChange={(val, idx) => setDepartment(val)}
      >
        {departments.map((val, idx) => {
          return <Picker.Item label={val} value={val} key={idx} />;
        })}
      </Picker>
    </View>
  );
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  picker: {
    height: 50,
    width: 200,
  },
});
