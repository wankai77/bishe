import { defineStore } from "pinia";

export const useUserStore = defineStore({
  id: "user",
  state: () => {
    return {
      token: "",
      username: "",
    };
  },
  actions: {
    setToken(token) {
      this.token = token;
    },
    clearToken() {
      this.token = "";
      this.username = "";
    },
    setUsername(name) {
      this.username = name;
    },
  },
  // 开启数据缓存
  persist: {
    enabled: true,
    strategies: [
      {
        storage: localStorage,
        paths: ["token", "username"],
      },
    ],
  },
});
