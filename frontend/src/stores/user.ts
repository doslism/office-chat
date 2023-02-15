import { ref } from "vue";
import { defineStore } from "pinia";

export const useCounterStore = defineStore("user", () => {
  const accessToken = ref("");
  const refreshToken = ref("");
  const username = ref("");
  function setUsername(name: string) {
    username.value = name;
  }
  function setAccessToken(token: string) {
    accessToken.value = token;
  }
  function setRefreshToken(token: string) {
    refreshToken.value = token;
  }

  return {
    username,
    accessToken,
    setUsername,
    setAccessToken,
    setRefreshToken,
  };
});
