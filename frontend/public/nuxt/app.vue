<template>
  <v-app>
    <v-app-bar title="KOシステム">
      <v-btn
        icon="mdi-github"
        href="https://github.com/ogoshikazuki/skill-sheet"
        target="_blank"
      />
    </v-app-bar>
    <v-navigation-drawer>
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
useHead({
  script: [
    {
      type: 'text/javascript',
      async: true,
      src: 'https://platform.twitter.com/widgets.js'
    }
  ]
})
</script>
