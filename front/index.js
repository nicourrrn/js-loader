

const load_elements = () => new Promise((res) =>{
  fetch("http://localhost:8080/get_products", {
    method: "GET"
  }).then(data => data.json().then(json_data => res(json_data)))
})

const main = async () => {
    let elements = await load_elements()
    let main_div = document.getElementById("product_list")
    for (let el of elements) {
        let prod_div = document.createElement("div")
        prod_div.innerText = `${el.name}: ${el.price}`
        main_div.appendChild(prod_div)
    }
}

const loop = () => {
    console.log("Еще один круг")
    new Promise((resolve) => {
        main()
        setTimeout(() => {
            resolve()
        }, 1000)
    })
}
loop()
