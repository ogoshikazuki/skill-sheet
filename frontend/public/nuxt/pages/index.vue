<template>
  <div>
    <h1>基本情報</h1>
    <ul>
      <li>生年月日: {{ data.basicInformation.birthday }}</li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import repositoryFactory from "~/repository/repositoryFactory"

const data = await repositoryFactory.get("basicInformation").find()
</script>
