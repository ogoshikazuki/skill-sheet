import { Breadcrumbs } from "./composables/states"

declare module '#app' {
  interface PageMeta {
    breadcrumbs: Breadcrumbs
  }
}
