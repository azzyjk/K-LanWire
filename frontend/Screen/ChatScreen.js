import React from "react";
import { Component } from "react";
import { View, Text, Dimensions, TouchableOpacity } from "react-native";
import { GiftedChat, Bubble } from "react-native-gifted-chat";
import axios from "axios";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { Ionicons } from "@expo/vector-icons";

const { height, width } = Dimensions.get("window");

class ChatScreen extends Component {
  constructor(props) {
    super(props);
    AsyncStorage.getItem("entranceYear", (err, result) => {
      this.entranceYear = result;
      AsyncStorage.getItem("department", (err, result) => {
        this.department = result;

        if (this.department) {
          if (this.department == "컴퓨터공학부") {
            this.departmentId = "c";
          } else if (this.department == "전기전자공학부") {
            this.departmentId = "e";
          } else if (this.department == "기계항공공학부") {
            this.departmentId = "m";
          } else {
            this.departmentId = "";
          }
          this.departmentId =
            this.departmentId + this.entranceYear[2] + this.entranceYear[3];
        }
      });
    });
    this.state = {
      log: [],
      msg: "",
      count: 1000,
    };
  }

  componentDidMount() {
    var danbee_head = {
      "Content-Type": "application/json;charset=UTF-8",
    };

    var danbee_body = {
      user_id: "cse17",
    };

    var danbee_url =
      "https://danbee.ai/chatflow/chatbot/v2.0/1d1e0598-3c44-4a8e-bd41-908a9c14846a/welcome.do";

    var options = {
      headers: danbee_head,
      data: JSON.stringify(danbee_body),
      method: "POST",
      url: danbee_url,
    };

    axios(options).then((res) => {
      var preLog = [];
      var tmpMessage = {
        _id: this.state.count,
        text: res.data.responseSet.result.result[0].message,
        createdAt: new Date(),
        user: {
          _id: 2,
          name: "React Native",
          avatar:
            "https://blog.kakaocdn.net/dn/esugIB/btqwv7PLOZT/oMb10ItMqbhdX0dKomtDQ0/img.jpg",
        },
      };
      preLog = [...preLog, tmpMessage];
      this.setState({
        log: preLog,
        msg: tmpMessage,
        count: this.state.count + 1,
        tmp: 0,
      });

      //console.log(this.entranceYear);
    });
  }

  renderBubble = (props) => {
    return (
      <Bubble
        {...props}
        textStyle={{
          right: {
            color: "white",
          },
          left: {
            color: "black",
          },
        }}
        wrapperStyle={{
          right: {
            backgroundColor: "#99CCFF",
          },
          left: {
            backgroundColor: "#CCFFCC",
          },
        }}
      />
    );
  };

  render() {
    const onSend = (messages) => {
      var danbee_head = {
        "Content-Type": "application/json;charset=UTF-8",
      };

      var danbee_body = {
        input_sentence: messages[0].text,
        user_id: this.departmentId,
      };

      var danbee_url =
        "https://danbee.ai/chatflow/chatbot/v2.0/1d1e0598-3c44-4a8e-bd41-908a9c14846a/message.do";

      var options = {
        headers: danbee_head,
        data: JSON.stringify(danbee_body),
        method: "POST",
        url: danbee_url,
      };

      axios(options).then((res) => {
        var preLog = this.state.log;

        var cnt = this.state.count + 1;
        var replyMsg = res.data.responseSet.result.result[0].message;

        var reply;
        var send;

        reply = {
          _id: cnt,
          text: replyMsg,
          createdAt: new Date(),
          user: {
            _id: 2,
            name: "React Native",
            avatar:
              "https://blog.kakaocdn.net/dn/esugIB/btqwv7PLOZT/oMb10ItMqbhdX0dKomtDQ0/img.jpg",
          },
        };
        preLog = [messages[0], ...preLog];
        preLog = [reply, ...preLog];

        if (res.data.responseSet.result.another_result[0]) {
          for (var i in res.data.responseSet.result.another_result) {
            send = {
              _id: cnt + 10000,
              text: (
                <Text
                  style={{
                    fontStyle: "italic",
                    fontWeight: "bold",
                    textDecorationLine: "underline",
                  }}
                  onPress={(event) => {
                    var tmp = [];
                    var rp = {
                      _id: cnt + 20000,
                      text: event._dispatchInstances.memoizedProps.children,
                      createdAt: new Date(),
                      user: {
                        _id: 1,
                        name: "React Native",
                        avatar:
                          "https://blog.kakaocdn.net/dn/esugIB/btqwv7PLOZT/oMb10ItMqbhdX0dKomtDQ0/img.jpg",
                      },
                    };
                    cnt++;
                    tmp = [...tmp, rp];
                    onSend(tmp);
                  }}
                >
                  {res.data.responseSet.result.another_result[i].intent_name}
                </Text>
              ),
              createdAt: new Date(),
              user: {
                _id: 2,
                name: "React Native",
                avatar:
                  "https://blog.kakaocdn.net/dn/esugIB/btqwv7PLOZT/oMb10ItMqbhdX0dKomtDQ0/img.jpg",
              },
            };
            cnt++;
            preLog = [send, ...preLog];
          }
        }

        this.setState({
          log: preLog,
          msg: messages[0],
          count: cnt,
        });
      });
    };

    return (
      <View style={{ flex: 1 }}>
        <View
          style={{
            flex: 1,
            backgroundColor: "#fff",
            flexDirection: "row",

            alignItems: "center",
            paddingTop: 20,
          }}
        >
          <TouchableOpacity
            style={{ flex: 1, paddingLeft: 10 }}
            onPress={() => this.props.navigation.navigate("Setting")}
          >
            <Ionicons name="md-settings" size={32} color="black" />
          </TouchableOpacity>
          <View
            style={{
              flex: 2,
              alignItems: "center",
            }}
          >
            <Text style={{ fontSize: 20, fontWeight: "900" }}>
              {this.entranceYear} {this.department}
            </Text>
          </View>
          <TouchableOpacity
            style={{ flex: 1, alignItems: "flex-end", paddingRight: 10 }}
            onPress={() => this.props.navigation.navigate("QnA")}
          >
            <Text style={{ fontSize: 33, fontWeight: "900" }}> Q</Text>
            {/* <Ionicons name="ios-people" size={32} color="black" /> */}
          </TouchableOpacity>
        </View>
        <View style={{ flex: 10 }}>
          <GiftedChat
            placeholder={"메세지를 입력하세요..."}
            alwaysShowSend={true}
            messages={this.state.log}
            renderBubble={this.renderBubble}
            textInputProps={{ keyboardAppearance: "dark", autoCorrect: false }}
            onSend={(messages) => onSend(messages)}
            user={{
              _id: 1,
            }}
          />
        </View>
      </View>
    );
  }
}

export default ChatScreen;
