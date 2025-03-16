document.addEventListener('DOMContentLoaded', () => {
    fetchItems();

    document.getElementById('itemForm').addEventListener('submit', async (e) => {
        e.preventDefault();

        const item = {
            id: document.getElementById('id').value,
            name: document.getElementById('name').value,
            price: parseFloat(document.getElementById('price').value)
        };

        try {
            const response = await fetch('/api/items', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(item)
            });

            if (!response.ok) {
                throw new Error(`Failed to add item: ${response.statusText}`);
            }

            const result = await response.json();
            document.getElementById('error').textContent = 'Item added successfully!';
            setTimeout(() => document.getElementById('error').textContent = '', 2000);
            fetchItems();
            e.target.reset();
        } catch (error) {
            document.getElementById('error').textContent = error.message;
        }
    });
});

async function fetchItems() {
    try {
        const response = await fetch('/api/items');
        if (!response.ok) {
            throw new Error(`Failed to fetch items: ${response.statusText}`);
        }

        const data = await response.json();
        console.log('Fetched data:', data);
        const itemList = document.getElementById('itemList');
        itemList.innerHTML = '';

        if (!data.items || Object.keys(data.items).length === 0) {
            itemList.innerHTML = '<li>No items available.</li>';
            return;
        }

        for (const [id, item] of Object.entries(data.items)) {
            const li = document.createElement('li');
            li.textContent = `ID: ${item.id}, Name: ${item.name}, Price: $${item.price.toFixed(2)}`;
            itemList.appendChild(li);
        }
    } catch (error) {
        console.error('Fetch error:', error);
        document.getElementById('error').textContent = error.message;
    }
}