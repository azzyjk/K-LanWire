import React from "react";
import { View, StyleSheet } from "react-native";
import { Picker } from "@react-native-picker/picker";

export default function EntranceYearPicker({
  entranceYear,
  setEntranceYear,
  job,
}) {
  const years = ["2019", "2018", "2017"];
  if (job == "Professor") return <View />;
  return (
    <View style={styles.container}>
      <Picker
        selectedValue={entranceYear}
        style={styles.picker}
        onValueChange={(val, idx) => setEntranceYear(val)}
      >
        {years.map((val, idx) => {
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
