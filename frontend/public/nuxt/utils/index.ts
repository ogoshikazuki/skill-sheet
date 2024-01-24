import { Gender } from '../graphql'

export const convertGenderForDisplay = (gender: Gender) => {
  switch (gender) {
    case Gender.Male:
      return '男性'
    case Gender.Female:
      return '女性'
  }
}
