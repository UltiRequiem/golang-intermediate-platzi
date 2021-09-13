// They like to complicate their lives xD
class Employee {
  private name: string;
  private id: number;

  public getName(): string {
    return this.name;
  }

  public setName(name: string) {
    this.name = name;
  }

  public getId(): number {
    return this.id;
  }

  public setId(id: number) {
    this.id = id;
  }
}
