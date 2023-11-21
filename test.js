import terminus from 'k6/x/terminus';
import { describe, expect } from 'https://jslib.k6.io/k6chaijs/4.3.4.3/index.js';

export default function () {
  describe('Terminus command execution', () => {
    const response = terminus.runTerminus("terminus whoami")
    console.log("Output is ====", response)
  });
}
