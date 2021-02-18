import React from "react";

import { NavigationContainer } from "@react-navigation/native";
import { createStackNavigator } from "@react-navigation/stack";

import SettingScreen from "../Screen/SettingScreen";
import QnAScreen from "../Screen/QnAScreen";
import QuestionScreen from "../Screen/QuestionScreen";
import ChatbotScreen from "../Screen/ChatScreen";

const Stack = createStackNavigator();

export default function Navigator() {
  return (
    <NavigationContainer>
      <Stack.Navigator screenOptions={{ headerShown: false }}>
        <Stack.Screen name="Chat" component={ChatbotScreen} />
        <Stack.Screen name="Setting" component={SettingScreen} />
        <Stack.Screen name="QnA" component={QnAScreen} />
        <Stack.Screen name="Question" component={QuestionScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
