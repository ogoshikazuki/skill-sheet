import { describe, expect, test } from 'vitest'
import { Gender } from '~/graphql'

describe('convertGenderForDisplay', () => {
  test('female', () => {
    const female = Gender.Female
    expect(convertGenderForDisplay(female)).toBe('女性')
  })

  test('male', () => {
    const male = Gender.Male
    expect(convertGenderForDisplay(male)).toBe('男性')
  })
})
