const purgecss = require('@fullhuman/postcss-purgecss')

module.exports = {
  plugins:
    (process.env.NODE_ENV === 'production')
      ?
      [
        require('postcss-csso'),
        purgecss({
          content: [`./index.html`, `./public/**/*.html`, `./src/**/*.vue`],
          defaultExtractor(content) {
            const contentWithoutStyleBlocks = content.replace(/<style[^]+?<\/style>/gi, '')
            return contentWithoutStyleBlocks.match(/[A-Za-z0-9-_/:]*[A-Za-z0-9-_/]+/g) || []
          },
          safelist: {
            standard: [/-(leave|enter|appear)(|-(to|from|active))$/, /^(?!(|.*?:)cursor-move).+-move$/, /^router-link(|-exact)-active$/, /data-v-.*/],
            deep: [/dark-theme/],
            greedy: []
          }
        })
      ]
      :
      [
        // No transformations in development
      ],
}