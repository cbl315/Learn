# -*-coding: utf8 -*-
# built-in package

# third-party package
import input_data
import tensorflow as tf

# custom package

"""
author: blcai
create: 11/5/18
data_sets: MinstHandWriteernDIgits
"""

LearnRate = 0.01
BatchRun = 1000


if __name__ == "__main__":
    # read input using input_data
    mnist = input_data.read_data_sets("MnistHandWriteernDigits", one_hot=True)
    # declare x as 28*28 matrix
    x =  tf.placeholder("float", [None, 784])
    # declare weight and bias
    w = tf.Variable(tf.zeros([784, 10]))
    b = tf.Variable(tf.zeros([10]))
    # define activation function y
    y = tf.nn.softmax(tf.matmul(x, w)+b)
    # declare labels
    y_ = tf.placeholder("float", [None, 10])
    # define loss
    cross_entropy = -tf.reduce_sum(y_*tf.log(y))
    # choose train optimization
    train_step = tf.train.GradientDescentOptimizer(LearnRate).minimize(cross_entropy)

    # start train
    init = tf.global_variables_initializer()
    sess = tf.Session()
    sess.run(init)
    for i in range(BatchRun):
        batch_xs, batch_ys = mnist.train.next_batch(100)
        sess.run(train_step, feed_dict={x: batch_xs, y_: batch_ys})

    # evaluate model perf
    correct_prediction = tf.equal(tf.argmax(y, 1), tf.argmax( y_, 1))
    accuracy = tf.reduce_mean(tf.cast(correct_prediction, "float"))
    print(sess.run(accuracy, feed_dict={x: mnist.test.images, y_:mnist.test.labels}))
