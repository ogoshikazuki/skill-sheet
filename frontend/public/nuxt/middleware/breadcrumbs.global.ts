import type { Breadcrumbs } from '~/breadcrumbs/types'

export default defineNuxtRouteMiddleware((to) => {
  const breadcrumbs = useBreadcrumbs()

  const toBreadcrumbs = to.meta.breadcrumbs as (Breadcrumbs | undefined)

  if (toBreadcrumbs === undefined) {
    throw createError('ページのパンくずリストが設定されていません。')
  }

  breadcrumbs.value = toBreadcrumbs
})
