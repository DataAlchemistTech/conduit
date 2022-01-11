import { RestSerializer } from 'ember-cli-mirage';

export default RestSerializer.extend({
  root: false,
  embed: true,

  normalize(payload) {
    const typeKey = this.typeKey;
    payload = { [typeKey]: payload };

    return RestSerializer.prototype.normalize.apply(this, [payload]);
  },
});
