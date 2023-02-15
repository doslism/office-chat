<template>
  <div class="register-form">
    <main>
      <div class="input-item">
        <div class="input-box">
          <label>Username</label>
          <input
            type="text"
            v-model="formData.username"
            @blur="onUsernameBlur"
            @focus="onUsernameFocus"
            :class="{ 'on-error': !isUsernameValid[0] }"
          />
        </div>
        <div class="input-error">{{ isUsernameValid[1] }}</div>
      </div>
      <div class="input-item">
        <div class="input-box">
          <label>Password</label>
          <input
            type="password"
            v-model="formData.password"
            @blur="onPasswordBlur"
            @focus="onPasswordFocus"
            :class="{ 'on-error': !isPasswordValid[0] }"
          />
        </div>
        <div class="input-error">{{ isPasswordValid[1] }}</div>
      </div>
      <div class="submit-btn" @click="submit">Submit</div>
    </main>
  </div>
</template>
<script lang="ts" setup>
import { reactive, ref } from "vue";
import axios from "axios";

const formData = reactive({
  username: "",
  password: "",
});

const fetching = ref(false);

const isUsernameValid = ref([true, ""]);

const onUsernameBlur = () => {
  validateUsername();
};

const onUsernameFocus = () => {
  isUsernameValid.value = [true, ""];
};

const validateUsername = () => {
  const val = formData.username;
  if (val.length < 6 || val.length > 16) {
    isUsernameValid.value[0] = false;
    isUsernameValid.value[1] = "Username length should be between 8 and 16.";
    isUsernameValid.value = [...isUsernameValid.value];
  }
};

const isPasswordValid = ref([true, ""]);

const onPasswordBlur = () => {
  validatePassword();
};

const onPasswordFocus = () => {
  isPasswordValid.value = [true, ""];
};

const validatePassword = () => {
  const val = formData.password;
  if (val.length < 6 || val.length > 16) {
    isPasswordValid.value[0] = false;
    isPasswordValid.value[1] = "Password length should be between 8 and 16.";
    isPasswordValid.value = [...isPasswordValid.value];
    return;
  }
  let reg = new RegExp(/^(?![^a-zA-Z]+$)(?!\D+$)/);
  if (!reg.test(val)) {
    isPasswordValid.value[0] = false;
    isPasswordValid.value[1] =
      "Should contain at least 1 character and 1 number.";
    isPasswordValid.value = [...isPasswordValid.value];
    return;
  }
};

const submit = () => {
  if (!fetching.value) {
    fetching.value = true;
    const xhr = axios.create();
    const payload = {
      username: formData.username,
      password: btoa(encodeURIComponent(formData.password)),
    };
    xhr
      .post("http://127.0.0.1:14000", payload)
      .then(() => {
        fetching.value = false;
      })
      .catch(() => {
        fetching.value = false;
      });
  }
};
</script>
<style scoped lang="scss">
.register-form {
  padding: 16px;

  main {
    margin: 0 auto;
    width: 500px;
  }

  .input-item {
    .input-box {
      display: flex;
      align-items: center;
      label {
        width: 100px;
        font-weight: 700;
      }
      input[type="text"],
      input[type="password"] {
        width: 300px;
        padding: 4px;
        outline: none;
      }

      input.on-error {
        border-width: 1px;
        background: rgb(247, 227, 227);
        border-color: rgb(250, 110, 110);
      }
    }
    .input-error {
      font-size: 12px;
      line-height: 14px;
      margin-left: 100px;
      color: rgb(250, 110, 110);
      padding: 4px;
    }
  }

  .submit-btn {
    margin-top: 12px;
    max-width: 80px;
    text-align: center;
    padding: 4px 8px;
    background: #f0f4f22a;
    font-weight: 900;
    cursor: pointer;
  }
}
</style>
