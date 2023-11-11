// This configuration is sourced from:
// https://strdr4605.com/commitlint-custom-commit-message-with-emojis

const matchAnyEmojiWithSpaceAfter =
  /(\u00a9|\u00ae|[\u2000-\u3300]|\ud83c[\ud000-\udfff]|\ud83d[\ud000-\udfff]|\ud83e[\ud000-\udfff])\s/
const matchTicketNumberWithSpaceAfter = /(?:\[(TEA-\d+)\]:\s)?/ // "[TEA-4605] ", "[TEA-1]"
const subjectThatDoesntStartWithBracket = /([^\[].+)/ // "Add tests" but don't allow "[ Add tests"

module.exports = {
  parserPreset: {
    parserOpts: {
      headerPattern: new RegExp(
        '^' +
          matchAnyEmojiWithSpaceAfter.source +
          matchTicketNumberWithSpaceAfter.source +
          subjectThatDoesntStartWithBracket.source +
          '$'
      ),
      headerCorrespondence: ['emoji', 'ticket', 'subject']
    }
  },
  plugins: [
    {
      rules: {
        'header-match-team-pattern': parsed => {
          const { emoji, ticket, subject } = parsed
          if (emoji === null && ticket === null && subject === null) {
            return [false, "header must be in format '✅ [TEA-123]: Description of work'"]
          }
          return [true, '']
        },
        'explained-emoji-enum': (parsed, _when, emojisObject) => {
          const { emoji } = parsed
          if (emoji && !Object.keys(emojisObject).includes(emoji)) {
            return [
              false,
              `emoji must be one of: ${Object.keys(emojisObject)
                .map(emojiType => `${emojiType} - ${emojisObject[emojiType]}`)
                .join('\n')}`
            ]
          }
          return [true, '']
        }
      }
    }
  ],
  rules: {
    'header-match-team-pattern': [2, 'always'],
    'explained-emoji-enum': [
      2,
      'always',
      {
        '✨': 'New feature',
        '🐞': 'Bug fix',
        '✅': 'Add or update tests',
        '🚧': 'Work in progress',
        '♻️': 'Refactor',
        '📝': 'Documentation update'
      }
    ]
  }
}
