import basicInformation from "./basicInformation";

interface Repositories {
  basicInformation: typeof basicInformation;
}

const repositories: Repositories = {
  basicInformation: basicInformation,
};

export default {
  get: (name: keyof Repositories) => repositories[name],
};
