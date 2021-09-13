// Very Cool, much Chad && Javascript FTW
class JSEmployee {
  constructor({ id, name }) {
    this._id = id;
    this._name = name;
  }

  doWork() {
    return `${this._name} is working.`;
  }

  get name() {
    return this._name.toUpperCase();
  }

  set name(name) {
    this.name = name;
  }

  get id() {
    return this._id;
  }

  set id(id) {
    this._id = id;
  }
}

const me = new JSEmployee({ id: 999, name: "Zero" });

console.log(me.name);
console.log(me.id);
console.log(me.doWork())
