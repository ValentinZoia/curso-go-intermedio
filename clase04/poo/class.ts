//Ejemplo de poo. En este caso con typeScript

class Camisetas {
  color1: string;
  color2: string;
  id: number;
  team: string;

  constructor(color1: string, color2: string, id: number, team: string) {
    this.color1 = color1;
    this.color2 = color2;
    this.id = id;
    this.team = team;
  }

  mostrarInfo(): string {
    return `ID: ${this.id}, Equipo: ${this.team}, Colores: ${this.color1}, ${this.color2}`;
  }

  cambiarEquipo(nuevoEquipo: string): void {
    this.team = nuevoEquipo;
  }
}
let miCamiseta = new Camisetas("Azul", "Rojo", 10, "San lorenzo");

console.log(miCamiseta.mostrarInfo());//ID: 10, Equipo: San lorenzo, Colores: Azul, Rojo
miCamiseta.cambiarEquipo("Boca");
console.log(miCamiseta.mostrarInfo());//ID: 10, Equipo: Boca, Colores: Azul, Rojo
