<template>
  <div>
    <h1>基本情報</h1>
    <ul>
      <li>生年月日: {{ data?.basicInformation.birthday }}</li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { useBasicInformationQuery } from '~/graphql'

const data = useBasicInformationQuery().result
</script>
