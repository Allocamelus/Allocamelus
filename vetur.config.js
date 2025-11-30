/** @type {import('vls').VeturConfig} */
module.exports = {
  // **optional** default: `{}`
  settings: {
    "vetur.useWorkspaceDependencies": true,
    "vetur.experimental.templateInterpolationService": true,
  },
  // **optional** default: `[{ root: './' }]`
  projects: ["./web/app"],
};
