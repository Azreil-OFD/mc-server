<template>
  <div class="min-h-screen flex flex-col">

    <!-- Меню для ПК -->
    <TabMenu v-if="!isMobile" :model="menuItems" class="p-1 shadow-md bg-white" />

    <!-- Контент страниц -->
    <div class="flex-grow overflow-auto p-4">
      <div v-if="activePage === 0">
        <Card class="shadow-md">
          <template #title>{{ username }}</template>
          <template #content>
            <div class="flex items-center">
              <Badge
                :severity="userOnlineStatus ? 'success' : 'danger'"
                :value="userOnlineStatus ? 'Online' : 'Offline'"
                class="mr-2"
              />
              <span>Последний вход: {{ lastLogin }}</span>
            </div>
          </template>
        </Card>
      </div>

      <div v-else-if="activePage === 1">
        <DataTable :value="onlinePlayers" class="shadow-md">
          <Column field="name" header="Игрок"></Column>
          <Column field="uniqueId" header="ID"></Column>
          <Column header="Статус">
            <template #body="slotProps">
              <Badge
                :value="slotProps.data.online ? 'Online' : 'Offline'"
                :severity="slotProps.data.online ? 'success' : 'danger'"
              />
            </template>
          </Column>
        </DataTable>
      </div>

      <div v-else-if="activePage === 2">
        <h1>В разработке</h1>
      </div>
    </div>

    <!-- Меню для мобильных устройств -->
    <TabMenu v-if="isMobile" :model="menuItems" class="p-1 shadow-md bg-white" />
  </div>
</template>

<script lang="ts" setup>
import { useBreakpoints } from '@vueuse/core';
const data = ref([])

onMounted(async () => {
  data.value = (await useFetch('/api/getOnline')).data.value
})

const breakpoints = useBreakpoints({ mobile: 640 });
const isMobile = breakpoints.smaller('mobile');

const username = useLocalStorage('username', '');


const activePage = ref(0);

const menuItems = [
  { label: 'Мой Профиль', icon: 'pi pi-user', command: () => activePage.value = 0 },
  { label: 'Игроки онлайн', icon: 'pi pi-users', command: () => activePage.value = 1 },
  { label: 'Настройки', icon: 'pi pi-cog', command: () => activePage.value = 2 },
];

const userOnlineStatus = computed(() => {
  const user = data.value.find(player => player.name === username.value);
  return user ? true : false;
});

const lastLogin = computed(() => {
  const user = data.value.find(player => player.name === username.value);
  return user ? new Date(user.firstPlayed).toLocaleString() : 'неизвестно';
});

const onlinePlayers = computed(() => data.value.filter(player => player.online));
</script>

<style>
html, body, #__nuxt {
  height: 100%;
  margin: 0;
  padding: 0;
}
</style>