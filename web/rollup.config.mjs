import { nodeResolve } from '@rollup/plugin-node-resolve'

export default {
  plugins: [nodeResolve()],
  onwarn: (warning, handler) => {
    if (warning.code === 'THIS_IS_UNDEFINED') { return; }
    handler(warning);
  }
}