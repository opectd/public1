using Confluent.Kafka;
using Microsoft.EntityFrameworkCore;
using StackExchange.Redis;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Redis Configuration
var redisConfig = ConfigurationOptions.Parse(builder.Configuration.GetConnectionString("Redis"));
builder.Services.AddSingleton<IConnectionMultiplexer>(ConnectionMultiplexer.Connect(redisConfig));

// Kafka Configuration
builder.Services.AddSingleton<IProducer<Null, string>>(sp => {
    var config = new ProducerConfig {
        BootstrapServers = builder.Configuration.GetConnectionString("Kafka")
    };
    return new ProducerBuilder<Null, string>(config).Build();
});

var app = builder.Build();

// Configure the HTTP request pipeline
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

// Endpoints
app.MapGet("/seckill/products", () => {
    // Get seckill products from Redis
    return Results.Ok(new { message = "Seckill products list" });
});

app.MapPost("/seckill/order", async (IProducer<Null, string> producer) => {
    // Process seckill order and publish to Kafka
    await producer.ProduceAsync("seckill_orders", new Message<Null, string> { Value = "Seckill order created" });
    return Results.Ok(new { message = "Seckill order created" });
});

app.Run();
