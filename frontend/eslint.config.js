import antfu from '@antfu/eslint-config'

export default await antfu({
  overrides: {
    vue: {
      'vue/custom-event-name-casing': ['error', 'kebab-case', {
        ignores: [],
      }],
    },
  },
})
