# GoCV - Visão Computacional usando Go e OpenCV 4

[GoCV](https://gocv.io/) é um package que fornece uma interface em GO para 
utilizarmos as versões mais recentes da biblioteca de visão computacional OpenCV.

O GoCV tem como objetivo tornar a linguagem Go um cliente de "primeira classe" 
compatível com os mais recentes desenvolvimentos no ecossistema OpenCV.

O GoCV suporta [CUDA](https://en.wikipedia.org/wiki/CUDA) para aceleração de 
hardware usando GPUs Nvidia, além disso, o GoCV tem suporte para [Intel OpenVINO](https://www.intel.com/content/www/us/en/developer/tools/openvino-toolkit/overview.html).

## O Projeto
A proposta desse projeto era realizar a classificação dos objetos e realizar a contagem dos objetos classificados.

O código utiliza de um classificador Haar Cascade, que é um algoritmo de aprendizado baseado em AdaBoost, esse algoritmos usa de caracteristicas para realizar a detectacção e classificação do objeto.
Além disso, utilizamos de um algoritmo para realizar a contagem de elementos.

