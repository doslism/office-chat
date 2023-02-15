<template>
  <div>
    <div class="input-box">
      <label for="">Username</label>
      <input type="text" v-model="formData.username" />
    </div>
    <div class="input-box">
      <label for="">Password</label>
      <input type="password" v-model="formData.password" />
    </div>
    <button @click="submit">Submit</button>
  </div>
</template>
<script lang="ts" setup>
import { watchEffect, reactive } from "vue";
import axios from "axios";

const formData = reactive({
  username: "",
  password: "",
});

watchEffect(() => {
  console.log(formData.username);
  console.log(formData.password);
});

const submit = () => {
  axios.post("http://127.0.0.1:14000/api/users/new", formData);
};
</script>
