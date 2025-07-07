# import cv2
import numpy as np
from sklearn.mixture import GaussianMixture
from scipy.spatial import ConvexHull
import matplotlib.pyplot as plt
from sklearn.preprocessing import StandardScaler
from scipy.optimize import minimize


"""

There are a total of 11 tries. In the first 4 / 5 we want to set up the foundational values for each of controls or extreme values. 

X Y //all of these over here are 100 uL each 
Y P
X P

Each of the above would create different RGB values


P //this one would be the sole color that is 200 uL

We would like to to change certain weights to move from exploration to exploitation. 


As a human, what would I do?

RED 
BLUE
YELLOW

PURPLE = RED + BLUE
GREEN = BLUE + YELLOW
ORANGE = RED + YELLOW

"""

# Base RGB values (normalized to 0-1)
COLOR_RGB = {
    "R": np.array([1.0, 0.0, 0.0]),
    "B": np.array([0.0, 0.0, 1.0]),
    "Y": np.array([1.0, 1.0, 0.0]),
}

# Secondary Colors
COLOR_RGB["P"] = 0.5 * (COLOR_RGB["R"] + COLOR_RGB["B"])  # PURPLE
COLOR_RGB["G"] = 0.5 * (COLOR_RGB["B"] + COLOR_RGB["Y"])  # GREEN
COLOR_RGB["O"] = 0.5 * (COLOR_RGB["R"] + COLOR_RGB["Y"])  # ORANGE


def mix_colors(color1, color2, ratio1=0.5, ratio2=0.5, noise_std=0.02):
    color = ratio1 * COLOR_RGB[color1] + ratio2 * COLOR_RGB[color2]
    noise = np.random.normal(0, noise_std, size=color.shape)
    return np.clip(color + noise, 0, 1)


# List of experimentas (tag, RGB)
experiments = []

# Pure colors (200 uL)
for c in ["R", "B", "Y"]:
    experiments.append((c, COLOR_RGB[c]))

# 1:1 binary mixtures (100 µL each)
experiments += [
    ("P", mix_colors("R", "B")),
    ("G", mix_colors("B", "Y")),
    ("O", mix_colors("R", "Y")),
    ("X+Y", mix_colors("R", "Y")),  # same as O
    ("Y+P", mix_colors("Y", "P")),
    ("X+P", mix_colors("R", "P")),
]

# Create training data
X = np.array([rgb for _, rgb in experiments])
labels = [name for name, _ in experiments]

np.random.seed(39)

# Random RBG Colors

n_samples = 500
X = np.random.randint(0, 256, (n_samples, 3))

# Normalize RGB values to range [0, 1]
scaler = StandardScaler()
X_scaled = scaler.fit_transform(X)

gmm = GaussianMixture(n_components=3, covariance_type="full", random_state=42)
gmm.fit(X_scaled)

predictions = gmm.predict(X_scaled)

# Visualize Clustering
fig = plt.figure(figsize=(6, 5))
ax = fig.add_subplot(111, projection="3d")
ax.scatter(X[:, 0], X[:, 1], X[:, 2], c=predictions, s=100)

for i, label in enumerate(labels):
    ax.text(X[i, 0], X[i, 1], X[i, 2], label, fontsize=10)

ax.set_xlabel("Red")
ax.set_ylabel("Green")
ax.set_zlabel("Blue")

ax.set_title("Color Mixing Clusters (GMM)")

# _____________________________________________

base_colors = np.array(
    [
        COLOR_RGB["R"],
        COLOR_RGB["Y"],
        COLOR_RGB["B"],
    ]
)

target = np.array([0.608, 0.631, 0.647])  # Example green


def loss(weights):
    prediction = np.dot(weights, base_colors)
    return np.linalg.norm(prediction - target)


constraints = {"type": "eq", "fun": lambda w: np.sum(w) - 1}
bounds = [(0, 1)] * 3
initial_guess = np.array([1 / 3, 1 / 3, 1 / 3])

result = minimize(loss, initial_guess, bounds=bounds, constraints=constraints)

best_mix = result.x  # proportions of R, Y, B


# Show result
if result.success:
    best_mix = result.x
    predicted_rgb = np.dot(best_mix, base_colors)
    print("✅ Optimization succeeded")
    print("Optimal mixing ratios [R, Y, B]:", best_mix)
    print("Predicted RGB:", predicted_rgb)
    print("Target RGB:", target)
    print("Color difference (L2 norm):", np.linalg.norm(predicted_rgb - target))
else:
    print("❌ Optimization failed:", result.message)

plt.show()
