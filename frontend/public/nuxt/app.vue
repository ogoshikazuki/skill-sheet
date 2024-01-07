<template>
  <div>
    <h1>基本情報</h1>
    <ul>
      <li>生年月日: {{ data?.basicInformation?.birthday }}</li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
// eslint-disable-next-line no-undef
const query = gql`
  query {
    basicInformation {
      birthday
    }
  }
`;

type BasicInformationResult = {
  basicInformation: {
    birthday: string;
  };
};

// eslint-disable-next-line no-undef
const { data } = await useAsyncQuery<BasicInformationResult>(query);
</script>
