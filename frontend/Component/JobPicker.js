import React from "react";
import { View, StyleSheet } from "react-native";
import { Picker } from "@react-native-picker/picker";

export default function JobPicker({ job, setJob }) {
  const jobs = ["Professor", "Student"];
  return (
    <View style={styles.container}>
      <Picker
        selectedValue={job}
        style={styles.picker}
        onValueChange={(val, idx) => setJob(val)}
      >
        {jobs.map((val, idx) => {
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
