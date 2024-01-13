import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: "../../../graphql/*.graphqls",
  documents: "./repository/**/*.graphql",
  generates: {
    "repository/types.ts": { plugins: ["typescript"] },
    "./repository/": {
      preset: "near-operation-file",
      presetConfig: {
        baseTypesPath: "types.ts",
      },
      plugins: ["typescript-operations", "typescript-vue-apollo"],
    },
  },
};

export default config;
