'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    return queryInterface.createTable('articles', {
        id: {
          type: Sequelize.BIGINT,
          allowNull: false,
          primaryKey: true,
          autoIncrement: true
      },
      title:{
          type: Sequelize.STRING,
          allowNull: true,
      },
      author:{
        type: Sequelize.STRING,
        allowNull: true,
    },
      genre:{
        type: Sequelize.STRING,
        allowNull: true,
      },
      image_url: {
        type: Sequelize.TEXT,
        allowNull: true
      },
      created: {
        type: Sequelize.DATE,
        allowNull: true
      },
      price: {
        type: Sequelize.DECIMAL(10,2),
        allowNull: false
      },
      body:{
        type: Sequelize.TEXT,
        allowNull: true,
      },
      created_at: {
        type: Sequelize.DATE,
        allowNull: false
      },
      updated_at: {
          type: Sequelize.DATE,
          allowNull: false
      },
      deleted_at: {
          type: Sequelize.DATE,
          allowNull: true
      }
    })
  },

  down: async (queryInterface, Sequelize) => {
    return queryInterface.dropTable('artists');
  }
};
