import React, { useState, useEffect } from "react";
import { View } from "react-native";
import AsyncStorage from "@react-native-async-storage/async-storage";

import InitScreen from "./Screen/InitScreen";

export default function App() {
  const [haveInfo, setHaveInfo] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    (async () => {
      if (
        (await AsyncStorage.getItem("department")) != null &&
        (await AsyncStorage.getItem("department")) != null &&
        (await AsyncStorage.getItem("department")) != null
      )
        setHaveInfo(true);
      setIsLoading(true);
    })();
  }, []);

  if (isLoading == false) return <View />;
  if (haveInfo == false)
    return <InitScreen haveInfo={haveInfo} setHaveInfo={setHaveInfo} />;
  else return <View />;
}
