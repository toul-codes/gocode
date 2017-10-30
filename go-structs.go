package main 

import (
    "fmt"
) 

type power struct { 
    attack int 
    defense int
} 

type location struct { 
    x float32 
    y float32
    z float32
}

type nonPlayerCharacter struct { 
    name string
    speed int
    hp int 
    power power
    loc location
} 

type attacker struct { 
    attackpower int 
    dmgbonus int 
}

type sword struct{
    attacker
    twohanded bool
} 

type gun struct { 
    attacker
    bulletsremaining int
} 

type chair struct { 
    attacker
    legsremaining int
}

func (loc location) String() string{ 
    return fmt.Sprintf("(%f, %f, %f)", loc.x, loc.y, loc.z)
} 

func (loc location) euclideanDistance(target location) float 64 {
    return math.Sqrt(
        (loc.x-target.x)*(loc.x-target.x) + (loc.y-target.y)*(loc.y-target.y) + (loc.z-target.z)*(loc.z-target.z))
    }

func ( npc nonPlayerCharacter) distanceTo(target nonPlayerCharacter) float 64 { 
    return npc.loc.euclideanDistance(target.loc)
} 




func (s sword) Wield() bool { 
    fmt.Println("You've wielded a sword!")
    return true
} 

func (g gun) Wield() bool { 
    fmt.Println("You've wielded a gun!") 
    return true
}

func (c chair) Wield() bool {
    fmt.Println("You've wielded a chair! Are you having a rough day?") 
    return true
}


func wielder(w weapon) bool{ 
    fmt.Println("Wielding...")
    return w.Wield()
}

type weapon interface { 
    Wield() bool
}

func main() { 
    fmt.Println("Structs...") 

    demon := nonPlayerCharacter{
        name: "Alfred",
        speed: 21,
        hp: 1000,
        power: power{attack: 75, defense: 50},
        loc: location{x: 1075.123, y: 521.123, z: 211.231}, 
    }
    
    fmt.Println(demon)

    anotherDemon := nonPlayerCharacter{
        name: "Beezlebub", 
        speed: 30, 
        hp: 5000,
        power: power{ attack: 10, defense: 10},
        loc: location{x: 32.03, y: 72.45, z: 65.231},
    }
    
    fmt.Println(anotherDemon)


    sword1 := sword{attacker: attacker{attackpower:1, dmgbonus: 5}, twohanded: true}
    gun1 := gun{attacker: attacker{attackpower: 10, dmgbonus: 20}, bulletsremaining: 11} 
    chair1 := chair{attacker: attacker{attackpower: 23, dmgbonus: 4}, legsremaining: 1}
    fmt.Printf("Weapons: sword: %v, gun:%v\n", sword1, gun1)
    sword1.Wield()
    
    wielder(sword1) 
    wielder(gun1) 
    wielder(chair1)




    fmt.Printf("Npc %v is %f units away from Npc %v\n", demon, demon. dastanceTo(anotherDemon), anotherDemon)
    

}






