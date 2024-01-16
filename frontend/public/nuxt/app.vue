<template>
  <v-app>
    <v-app-bar>
      <v-app-bar-nav-icon @click="drawer = !drawer" />
      <v-app-bar-title>KOシステム</v-app-bar-title>
      <v-btn
        icon="mdi-github"
        href="https://github.com/ogoshikazuki/skill-sheet"
        target="_blank"
      />
    </v-app-bar>
    <v-navigation-drawer v-model="drawer">
      <v-list>
        <v-list-item>
          <a
            class="twitter-timeline"
            href="https://twitter.com/k_o_system?ref_src=twsrc%5Etfw"
          >Tweets by k_o_system</a>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <NuxtPage />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { useDisplay } from 'vuetify/lib/framework.mjs'

useHead({
  script: [
    {
      async: true,
      src: 'https://platform.twitter.com/widgets.js'
    }
  ]
})

const { mobile } = useDisplay()
const drawer = ref<boolean>(!mobile.value)
</script>
