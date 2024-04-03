    import * as THREE from 'three';
    import { OrbitControls } from 'three/addons/controls/OrbitControls.js';

    var rooms = document.getElementsByClassName("room")
    var relations = document.getElementsByClassName("relation")
    var start = document.querySelector(".EndPointStart").textContent
    var end = document.querySelector(".EndPointEnd").textContent
    const size = 4

    const renderer = new THREE.WebGLRenderer();

    renderer.setSize(window.innerWidth, window.innerHeight);

    document.body.appendChild(renderer.domElement);

    const scene = new THREE.PerspectiveCamera();

    const camera = new THREE.PerspectiveCamera(
        45,
        window.innerWidth / window.innerHeight,
        0.1,
        1000
    )

    const orbit = new OrbitControls(camera, renderer.domElement);

    const axesHelper = new THREE.AxesHelper(5);
    scene.add(axesHelper);

    camera.position.set(0, 30, 200);
    orbit.update();

    const boxGeometry = new THREE.BoxGeometry();
    const boxMaterial = new THREE.MeshBasicMaterial({color: 0x00FF00});
    const box = new THREE.Mesh(boxGeometry, boxMaterial);
    scene.add(box);

    const planeGeometry = new THREE.PlaneGeometry(100, 100);
    const planeMaterial = new THREE.MeshBasicMaterial({
        color: 0xFFFFFF,
        side: THREE.DoubleSide
    });
    const plane = new THREE.Mesh(planeGeometry, planeMaterial);
    scene.add(plane);
    plane.rotation.x = -0.5 * Math.PI;;
    plane.position.y = -30;

    const gridHelper = new THREE.GridHelper(100);
    scene.add(gridHelper);
    gridHelper.position.y = -30;

    for (let i = 0; i < rooms.length; i++) {
        const room = rooms[i];
        let name = room.querySelector(".name").textContent;
        const sphereGeometry = new THREE.SphereGeometry(size, 16, 16);
        const sphereMaterial = new THREE.MeshBasicMaterial({
            wireframe: true,
        });
        if (name == start) {
            sphereMaterial.color = new THREE.Color(0x00FF00);
        } else if (name == end) {
            sphereMaterial.color = new THREE.Color(0xFF0000);
        } else {
            sphereMaterial.color = new THREE.Color(0x0000FF);
        };
        
        const sphere = new THREE.Mesh(sphereGeometry, sphereMaterial);
        sphere.name = name
        let x = room.querySelector(".x").textContent;
        let y = room.querySelector(".y").textContent;
        sphere.position.set((x-50), 100-y-30+size, (Math.random()*100)-50);
        scene.add(sphere);
    }

    for (let i = 0; i < relations.length; i++) {
        const relation = relations[i]
        var roomOne = relation.querySelector(".firstRoom").textContent
        var roomTwo = relation.querySelector(".secondRoom").textContent
        var sphereOne = scene.getObjectByName(roomOne)
        var sphereTwo = scene.getObjectByName(roomTwo)
        scene.getObjectByName(roomTwo)

        var midPoint = new THREE.Vector3().addVectors(sphereOne.position, sphereTwo.position).multiplyScalar(0.5);
        var directionVector = new THREE.Vector3().subVectors(sphereTwo.position, sphereOne.position);

        const geometry = new THREE.CylinderGeometry(3, 3, directionVector.length(), 10, 30)
        const material = new THREE.MeshBasicMaterial({
            color: 0x00FFFF,
            wireframe: true
        });
        const cylinder = new THREE.Mesh(geometry, material);
        cylinder.position.copy(midPoint);

        // Calcule la rotation de l'axe z
        var quaternion = new THREE.Quaternion();
        quaternion.setFromUnitVectors(new THREE.Vector3(0, 1, 0), directionVector.clone().normalize());
        cylinder.quaternion.copy(quaternion);
        scene.add(cylinder);
    }
    
    function animate(time) {
        renderer.render(scene, camera)
    }

    renderer.setAnimationLoop(animate);

    renderer.render(scene, camera);
