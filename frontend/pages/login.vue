<template>
  <div class="flex items-center justify-center min-h-screen" style="background-image: url(https://cdn.mcgolem.com/ad2584f29d40f6f7c998588ece91d7fa.png); background-size:cover ">
    <div class="bg-white p-8 rounded-2xl shadow-xl w-full max-w-md" 
     >
      <h2 class="text-2xl font-semibold text-center text-gray-800 mb-6">Авторизация</h2>

      <Message
        severity="error"
        variant="outlined"
        class="mb-4"
        v-if="error"
      >
        {{ error }}
      </Message>

      <div class="space-y-4">
        <InputText
          type="text"
          placeholder="Логин"
          v-model="loginInput"
          class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
        />

        <InputText
          type="password"
          placeholder="Пароль"
          v-model="passwordInput"
          class="w-full border border-gray-300 rounded-lg px-4 py-3 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
        />
      </div>

      <Button
        label="Войти"
        class="mt-6 w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold rounded-lg py-3 transition duration-300"
        @click="loginProcess"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import Cookies from 'js-cookie';

const token = useLocalStorage("token", "");
const username = useLocalStorage("username", "");
const error = ref("");

const loginInput = ref("");
const passwordInput = ref("");

const loginProcess = async () => {
  error.value = "";
  const result = await useFetch('/api/login', {
    method: 'post',
    body: {
      login: loginInput.value,
      password: passwordInput.value
    }
  });

  // @ts-ignore
  if (result.data.value.status) {
    token.value = result.data.value.data.token;
    username.value = result.data.value.data.username;
    Cookies.set('token', result.data.value.data.token, { expires: 7 });

    navigateTo('/');
  } else {
    error.value = result.data.value.error.message;
  }
};
</script>


<style scoped>

</style>