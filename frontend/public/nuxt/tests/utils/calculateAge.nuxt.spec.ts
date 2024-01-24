import { Temporal } from 'temporal-polyfill'
import { describe, expect, test } from 'vitest'

describe('calculateAge', () => {
  test('誕生日当日', () => {
    const birthday = Temporal.PlainDate.from('1991-07-01')
    const now = Temporal.PlainDate.from('2023-07-01')
    expect(calculateAge(birthday, now)).toBe(32)
  })

  test('誕生日前日', () => {
    const birthday = Temporal.PlainDate.from('1991-07-01')
    const now = Temporal.PlainDate.from('2023-06-30')
    expect(calculateAge(birthday, now)).toBe(31)
  })
})
