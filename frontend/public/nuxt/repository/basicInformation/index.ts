import { useBasicInformationQuery } from "./basicInformation.generated";

export default {
  find() {
    return useBasicInformationQuery().result;
  },
};
