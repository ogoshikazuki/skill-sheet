<script lang="ts" setup>
import { useBasicInformationQuery } from '~/graphql'

const basicInformationQuery = useBasicInformationQuery().result
const basicInformations = [
  {
    title: '生年月日',
    value: basicInformationQuery.value?.basicInformation.birthday
  }
]
</script>

<template>
  <v-card>
    <v-list>
      <v-list-subheader>
        基本情報
      </v-list-subheader>
      <v-list-item
        v-for="basicInformation in basicInformations"
        :key="basicInformation.title"
        :title="basicInformation.title"
        :subtitle="basicInformation.value"
      />
    </v-list>
  </v-card>
</template>
