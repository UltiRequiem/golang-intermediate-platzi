// They like to complicate their lives xD
// Javascript FTW
class Employee {
  private employeeNick: string;
  private id: number;

  constructor(id: number, nick: string) {
    this.employeeNick = nick;
    this.id = id;
  }

  public getEmployeeNick(): string {
    return this.employeeNick;
  }

  public setEmployeeNick(nick: string) {
    this.employeeNick = nick;
  }

  public getId(): number {
    return this.id;
  }

  public setId(id: number) {
    this.id = id;
  }
}

const myEmployee = new Employee(999, "Default");

console.log(myEmployee.getId());
console.log(myEmployee.getEmployeeNick());

// Uṕdate Properties

myEmployee.setEmployeeNick("Zero");
myEmployee.setId(7777);

console.log(myEmployee.getId());
console.log(myEmployee.getEmployeeNick());
