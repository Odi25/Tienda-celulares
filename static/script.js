let usuario_id = null;
let rol = null;

async function login() {
  const correo = document.getElementById("correo").value;
  const password = document.getElementById("password").value;

  const res = await fetch("/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ correo, password })
  });

  const data = await res.json();

  if (res.ok) {
    rol = data.rol;
    usuario_id = data.id;

    document.getElementById("login-register-panel").style.display = "none";

    if (rol === "admin") {
      document.getElementById("admin-panel").classList.remove("oculto");
      cargarProductosAdmin(); // ✅ solo si es admin
    } else {
      document.getElementById("cliente-panel").classList.remove("oculto");
      cargarProductos(); // ✅ solo si es cliente
      cargarHistorialCompras(); // 🔄 Actualiza la tabla automáticamente
    }

    document.getElementById("login-msg").innerText = "";
  } else {
    document.getElementById("login-msg").innerText = "❌ " + (data.error || "Error de inicio de sesión");
  }

  document.getElementById("logout-btn").classList.remove("oculto");
}

async function cargarProductos() {
  const res = await fetch("/api/productos");
  const productos = await res.json();

  const contenedor = document.getElementById("productos-cards-cliente");
  if (!contenedor) {
    console.warn("❌ No se encontró el contenedor de tarjetas: #productos-cards-cliente");
    return;
  }

  contenedor.innerHTML = "";

  productos.forEach((p, i) => {
    const card = document.createElement("div");
    card.className = "col-md-4 mb-4";

    const imagen = p.imagen_url || "https://m.media-amazon.com/images/I/51nuxoc-u0L._AC_UF894,1000_QL80_.jpg";
    
    const producto = {
    nombre: document.getElementById("nombre").value,
    descripcion: document.getElementById("descripcion").value,
    precio: parseFloat(document.getElementById("precio").value),
    stock: parseInt(document.getElementById("stock").value),
    imagen_url: document.getElementById("imagen_url").value // ✅
    };
    document.getElementById("imagen_url").value = p.imagen_url || "";
    const disponible = p.stock > 0
      ? `<button class="btn btn-primary w-100" onclick="comprarProducto(${p.id})">Comprar</button>`
      : `<span class="badge bg-secondary w-100 py-2">Agotado</span>`;

    card.innerHTML = `
      <div class="card h-100 shadow">
        <img src="${imagen}" class="card-img-top p-3" alt="Celular" style="height: 200px; object-fit: contain;">
        <div class="card-body">
          <h5 class="card-title">${p.nombre}</h5>
          <p class="card-text">${p.descripcion}</p>
          <p class="card-text"><strong>Precio:</strong> $${p.precio}</p>
          <p class="card-text"><strong>Stock:</strong> ${p.stock}</p>
        </div>
        <div class="card-footer text-center">
          ${disponible}
        </div>
      </div>
    `;

    contenedor.appendChild(card);
  });
}

async function comprarProducto(producto_id) {
  if (!usuario_id) {
    alert("Debes iniciar sesión para comprar.");
    return;
  }

  const cantidad = parseInt(prompt("¿Cuántas unidades deseas comprar?", "1"));
  if (!cantidad || cantidad <= 0) return;

  const res = await fetch("/api/compras", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ usuario_id, producto_id, cantidad })
  });

  const data = await res.json();
  if (res.ok) {
    alert("✅ Compra registrada con éxito");
    cargarProductos();           // 🔄 Actualiza stock
    cargarHistorialCompras();    // 🔄 Actualiza historial
  } else {
    alert("❌ Error: " + (data.error || "No se pudo completar la compra"));
  }
}

async function cargarHistorialCompras() {
  const res = await fetch(`/api/compras/${usuario_id}`);
  const compras = await res.json();
  const tbody = document.querySelector("#tabla-compras tbody");
  tbody.innerHTML = "";

  for (const compra of compras) {
    // Obtener el nombre del producto
    let productoNombre = "ID: " + compra.producto_id;
    const productoRes = await fetch(`/api/productos/${compra.producto_id}`);
    if (productoRes.ok) {
      const producto = await productoRes.json();
      productoNombre = producto.nombre;
    }

    tbody.innerHTML += `
      <tr>
        <td>${compra.id}</td>
        <td>${productoNombre}</td>
        <td>${compra.cantidad}</td>
        <td>${new Date(compra.fecha).toLocaleString()}</td>
      </tr>
    `;
  }
}
async function cargarProductosAdmin() {
  const res = await fetch("/api/productos");
  const productos = await res.json();

  console.log("✅ Productos recibidos:", productos); // ✅ Verifica en consola

  const tbody = document.querySelector("#tabla-productos tbody");
  console.log("🔍 tbody existe:", tbody); // ✅ Verifica en consola

  tbody.innerHTML = "";

  productos.forEach(p => {
    tbody.innerHTML += `
      <tr>
        <td>${p.nombre}</td>
        <td>${p.descripcion}</td>
        <td>$${p.precio}</td>
        <td>${p.stock}</td>
        <td>
          <button onclick="editarProducto(${p.id})">✏️</button>
          <button onclick="eliminarProducto(${p.id})">🗑️</button>
        </td>
      </tr>`;
  });
}

document.getElementById("form-producto").addEventListener("submit", async e => {
  e.preventDefault();
  const id = document.getElementById("producto-id").value;

  const producto = {
    nombre: document.getElementById("nombre").value,
    descripcion: document.getElementById("descripcion").value,
    precio: parseFloat(document.getElementById("precio").value),
    stock: parseInt(document.getElementById("stock").value),
    imagen_url: document.getElementById("imagen_url").value  // ✅ Agregado
  };

  const url = id ? `/api/productos/${id}` : "/api/productos";
  const method = id ? "PUT" : "POST";

  await fetch(url, {
    method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(producto)
  });

  document.getElementById("form-producto").reset();
  cargarProductosAdmin();
});

function editarProducto(id) {
  fetch(`/api/productos/${id}`)
    .then(res => res.json())
    .then(p => {
      document.getElementById("producto-id").value = p.id;
      document.getElementById("nombre").value = p.nombre;
      document.getElementById("descripcion").value = p.descripcion;
      document.getElementById("precio").value = p.precio;
      document.getElementById("stock").value = p.stock;
      document.getElementById("imagen_url").value = p.imagen_url || ""; // ✅ Agregado
    });
}

function eliminarProducto(id) {
  if (confirm("¿Eliminar este producto?")) {
    fetch(`/api/productos/${id}`, { method: "DELETE" })
      .then(() => cargarProductosAdmin());
  }
}

function mostrarLogin() {
  document.getElementById("login-form").classList.remove("oculto");
  document.getElementById("registro-form").classList.add("oculto");
}

function mostrarRegistro() {
  document.getElementById("registro-form").classList.remove("oculto");
  document.getElementById("login-form").classList.add("oculto");
}

async function registrar() {
  const nombre = document.getElementById("nombre-reg").value;
  const correo = document.getElementById("correo-reg").value;
  const password = document.getElementById("password-reg").value;

  const res = await fetch("/api/registro", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nombre, correo, password })
  });

  const data = await res.json();

  if (res.ok) {
    document.getElementById("registro-msg").innerText = "✅ Registro exitoso. Ahora inicia sesión.";
    mostrarLogin();
  } else {
    document.getElementById("registro-msg").innerText = "❌ " + (data.error || "No se pudo registrar.");
  }
}

function logout() {
  location.reload(); // simplemente reinicia la página
}
