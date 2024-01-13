export default {
  async find() {
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
    const result = await useAsyncQuery<BasicInformationResult>(query);

    return result.data;
  },
};
