<script setup lang="ts">
onMounted(() => {
  useHead({
    script: [
      {
        async: true,
        src: 'https://platform.twitter.com/widgets.js'
      }
    ]
  })
})
</script>

<template>
  <a class="twitter-timeline" href="https://twitter.com/k_o_system?ref_src=twsrc%5Etfw">Tweets by
    k_o_system</a>
</template>
