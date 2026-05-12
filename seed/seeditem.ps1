$items = @(
    @{
        name = "Indomie Goreng"
        sku = "IND-001"
        price_sell = 3500
        price_buy = 2800
    },
    @{
        name = "Aqua 600ml"
        sku = "AQU-001"
        price_sell = 4000
        price_buy = 2500
    },
    @{
        name = "Pocari Sweat"
        sku = "POC-001"
        price_sell = 8000
        price_buy = 6000
    },
    @{
        name = "Teh Botol"
        sku = "TEH-001"
        price_sell = 5000
        price_buy = 3500
    }
)

foreach ($item in $items) {
    Invoke-RestMethod `
        -Uri "http://localhost:3000/api/v1/items" `
        -Method POST `
        -ContentType "application/json" `
        -Body ($item | ConvertTo-Json)
}