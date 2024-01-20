import { test, expect } from '@playwright/test'

test('basic information', async ({ page }) => {
  await page.goto('/basic-information')

  await page.getByText('1991-07-01').waitFor()
  await expect(page.getByText('1991-07-01')).toBeVisible()
})
